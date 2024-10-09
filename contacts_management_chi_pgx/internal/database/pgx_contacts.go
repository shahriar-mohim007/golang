package repository

import (
	"chi_pgx/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"strings"
	"sync"
)

type PgxRepository struct {
	db *pgxpool.Pool
}

var (
	once       sync.Once
	repository *PgxRepository
)

func NewPgRepository(databaseUrl string) (*PgxRepository, error) {
	var err error
	once.Do(func() {
		db, dbErr := pgxpool.New(context.Background(), databaseUrl)
		if dbErr != nil {
			err = dbErr
			log.Error().Err(dbErr).Msgf("Database Connection Error: %v", err)
			return
		}

		if pingErr := db.Ping(context.Background()); pingErr != nil {
			err = pingErr
			log.Error().Err(pingErr).Msgf("Database Ping Error: %v:", err)
			return
		}

		repository = &PgxRepository{db: db}
	})
	return repository, err
}

func (repo *PgxRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password, is_active FROM users WHERE email = $1`
	err := repo.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PgxRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password, is_active,created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	err := repo.db.QueryRow(ctx, query, user.ID, user.Name, user.Email, user.Password, user.IsActive).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgxRepository) ActivateUserByID(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE users SET is_active = TRUE WHERE id = $1`
	_, err := repo.db.Exec(ctx, query, userID)
	return err
}

func (repo *PgxRepository) GetAllContacts(ctx context.Context, userID uuid.UUID) ([]domain.Contact, error) {
	rows, err := repo.db.Query(ctx, "SELECT id, phone, street, city, state, zip_code, country FROM contacts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []domain.Contact
	for rows.Next() {
		var contact domain.Contact
		err := rows.Scan(&contact.ID, &contact.Phone, &contact.Street, &contact.City, &contact.State, &contact.ZipCode, &contact.Country)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, rows.Err()
}

func (repo *PgxRepository) CreateContact(ctx context.Context, contact *domain.Contact) error {
	query := `
        INSERT INTO contacts 
        (id, user_id, phone, street, city, state, zip_code, country, created_at, updated_at) 
        VALUES 
        ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
    `
	_, err := repo.db.Exec(
		ctx, query,
		contact.ID, contact.UserID, contact.Phone, contact.Street, contact.City, contact.State, contact.ZipCode, contact.Country,
	)
	return err
}

func (repo *PgxRepository) GetContactByID(ctx context.Context, contactID uuid.UUID) (*domain.ContactWithUserResponse, error) {
	query := `
        SELECT 
            contacts.id AS contact_id,
            contacts.phone,
            contacts.street,
            contacts.city,
            contacts.state,
            contacts.zip_code,
            contacts.country,
            users.name AS user_name,
            users.email AS user_email
        FROM 
            contacts
        JOIN 
            users ON contacts.user_id = users.id
        WHERE 
            contacts.id = $1;
    `

	var response domain.ContactWithUserResponse
	err := repo.db.QueryRow(ctx, query, contactID).Scan(
		&response.ContactID,
		&response.Phone,
		&response.Street,
		&response.City,
		&response.State,
		&response.ZipCode,
		&response.Country,
		&response.UserName,
		&response.UserEmail,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no contact found with ID: %s", contactID)
		}
		return nil, err
	}
	return &response, nil
}

func (repo *PgxRepository) PatchContact(ctx context.Context, contactID uuid.UUID, contact *domain.Contact) error {
	var queryParts []string
	var args []interface{}
	argID := 1

	if contact.Phone != "" {
		queryParts = append(queryParts, fmt.Sprintf("phone = $%d", argID))
		args = append(args, contact.Phone)
		argID++
	}
	if contact.Street != "" {
		queryParts = append(queryParts, fmt.Sprintf("street = $%d", argID))
		args = append(args, contact.Street)
		argID++
	}
	if contact.City != "" {
		queryParts = append(queryParts, fmt.Sprintf("city = $%d", argID))
		args = append(args, contact.City)
		argID++
	}
	if contact.State != "" {
		queryParts = append(queryParts, fmt.Sprintf("state = $%d", argID))
		args = append(args, contact.State)
		argID++
	}
	if contact.ZipCode != "" {
		queryParts = append(queryParts, fmt.Sprintf("zip_code = $%d", argID))
		args = append(args, contact.ZipCode)
		argID++
	}
	if contact.Country != "" {
		queryParts = append(queryParts, fmt.Sprintf("country = $%d", argID))
		args = append(args, contact.Country)
		argID++
	}

	if len(queryParts) == 0 {
		return fmt.Errorf("no fields provided to update")
	}

	query := fmt.Sprintf("UPDATE contacts SET %s WHERE id = $%d", strings.Join(queryParts, ", "), argID)
	args = append(args, contactID)

	_, err := repo.db.Exec(ctx, query, args...)
	return err
}

func (repo *PgxRepository) DeleteContactByID(ctx context.Context, contactID uuid.UUID) error {
	query := `DELETE FROM contacts WHERE id = $1`
	result, err := repo.db.Exec(ctx, query, contactID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
