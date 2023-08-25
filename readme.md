# Envcrypt

This project is a simple tool to encrypt and decrypt your .env files with your primary gpg key. This assumes that your .env file is called `.env` and that it is in the gitignore file. The encrypted file will be called `.env.asc` and should be committed to the repository.

## Usage

After any changes to the `.env` file, run `envcrypt` to encrypt the file. This will create a `.env.asc` file that should be committed to the repository. When you pull the repository, run `envcrypt` to decrypt the file.
