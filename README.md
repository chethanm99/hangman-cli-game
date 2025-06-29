# Hangman-CLI-Game

This is a simple classical CLI-based Hangman game written in Golang, which will ask users to guess a word, and for every wrong guess, the attempts will be reduced. You don't have to be a developer to run this game, all you need to do is follow the simple steps, then you're all set to play!

# Key Features

-Users can choose to play the game on different levels: Easy, Medium, and Hard.

-The Easy level will give 10 attempts to the user to guess the word. The Medium level will give 6 attempts, and the Hard level will only give 4 attempts to guess the word.

-On every wrong guess, the wrong letter will be displayed, which will help the user not to repeat the same letter.

-For every wrong attempt, the formation of hangman starts, and as soon as all the attempts are over, a full hangman will be formed, and the game will be finished.

-Finally, the summary of the user's gameplay (games played , wins , losses) will be displayed .

# How to play

Follow these steps to download and play the game directly.

### 1. Go to the Releases Page

Click the here to go to the official downloads page for this game : https://github.com/chethanm99/hangman-cli-game/releases/tag/v1.0.0

### 2. Download the Correct Files

On the releases page, look for the "Assets" section. You need to download two files:

1.  **The Game Executable:**
    *   If you are on **Windows**, download `hangman.exe`.
    *   If you are on **macOS**, download `hangman-macos`.
    *   If you are on **Linux**, download `hangman-linux`.
  
2.  **The Word List:**
    *   You **must** also download the `words.txt` file.

### 3. Put the Files in the Same Folder

Create a new folder on your computer (e.g., on your Desktop) and place **both** the game executable and the `words.txt` file inside it.

### 4. Run the Game!

*   **On Windows:** Simply double-click the `hangman.exe` file.
*   **On macOS or Linux:** You need to give the file permission to run first.
  
    1. Open your Terminal application.
    2. Navigate to the folder where you saved the files (e.g., `cd Desktop/MyHangmanGame`).
    3. Run the command: `chmod +x ./hangman-macos` (or `./hangman-linux`).
    4. Now, run the game with: `./hangman-macos` (or `./hangman-linux`).
 
---

## For Developers (Building from Source)

If you have Go installed and want to build the project yourself:

1.  **Clone the Repository**
    ```bash
    git clone https://github.com/chethanm99/your-repo-name.git
    cd your-repo-name
    ```

2.  **Create the Word File**
    Make sure you have a `words.txt` file in the root directory with one word per line.

3.  **Run the Game**
    ```bash
    go run .
    ```
