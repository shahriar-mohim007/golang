import psycopg2
import string
import random

def connect_db():
    return psycopg2.connect(
        host="localhost",
        database="url_shortener",
        user="postgres",
        password="postgres"
    )


BASE62 = string.ascii_letters + string.digits


def generate_short_url(length=6):
    return ''.join(random.choice(BASE62) for _ in range(length))

def insert_url(original_url, short_url):
    conn = connect_db()
    cursor = conn.cursor()
    cursor.execute("INSERT INTO urls (original_url, short_url) VALUES (%s, %s) RETURNING id;", (original_url, short_url))
    conn.commit()
    cursor.close()
    conn.close()


def get_original_url(short_url):
    conn = connect_db()
    cursor = conn.cursor()
    cursor.execute("SELECT original_url FROM urls WHERE short_url = %s;", (short_url,))
    result = cursor.fetchone()
    cursor.close()
    conn.close()
    return result[0] if result else None


def shorten_url(original_url):
    short_url = generate_short_url()
    insert_url(original_url, short_url)
    return short_url


if __name__ == '__main__':
    original_url = "https://www.example.com"
    

    short_url = shorten_url(original_url)
    print(f"Short URL: {short_url}")
    
    
    retrieved_url = get_original_url(short_url)
    print(f"Original URL: {retrieved_url}")
