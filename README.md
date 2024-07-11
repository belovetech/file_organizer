# file_organizer

File Organizer is a command-line tool that organizes files in a specified directory based on their file extensions. The program creates subdirectories named after the file extensions and moves the files into their respective subdirectories. Files without an extension are moved into a directory named no_extension.

## Features

- Organizes files by their extensions.
- Creates subdirectories named after the extensions.
- Moves files without an extension into a no_extension directory.
- Supports optional flags for verbose output and dry-run mode.

## Usage

To use the File Organizer, pass the path of the directory you want to organize as an argument to the program.

```sh
./file-organizer --dir=../testDir
```

## Optional Flags

`--verbose`: Enables verbose output to show detailed information about the organization process.

`--dry-run`: Performs a dry run without actually moving any files. Useful for seeing what changes will be made.

### Examples

##### Organize Files with Verbose Output

```sh
./file-organizer --dir=../testDir --verbose
```

#### Perform a Dry Run

```sh
./file-organizer --dir=../testDir --dry-run
```

### Installation

To build the File Organizer from source, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/yourusername/file-organizer.git
cd file-organizer
```

2. Build the project:

```sh
go build -o file-organizer
```

3. Run the executable with the desired flags and directory path.
