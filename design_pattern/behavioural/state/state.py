class State:
    def handle(self, context):
        pass

class PlayingState(State):
    def handle(self, context):
        print("Currently playing. Transitioning to paused state.")
        context.state = PausedState()

class PausedState(State):
    def handle(self, context):
        print("Currently paused. Transitioning to playing state.")
        context.state = PlayingState()

class MediaPlayer:
    def __init__(self, state):
        self.state = state

    def press_button(self):
        self.state.handle(self)

# Usage
player = MediaPlayer(PlayingState())
player.press_button()  # Switch to paused
player.press_button()  # Switch to playing
