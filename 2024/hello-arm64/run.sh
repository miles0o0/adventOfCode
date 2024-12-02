#!/bin/bash

# Function to display usage instructions
usage() {
    echo "Usage: $0 -d <day_directory>"
    echo "Example: $0 -d hello-arm64"
    exit 1
}

# Parse command-line arguments
while getopts ":d:" opt; do
    case $opt in
        d)
            DAY=$OPTARG
            ;;
        *)
            usage
            ;;
    esac
done

# Check if -d flag was provided
if [ -z "$DAY" ]; then
    usage
fi

# Define the directory and file paths
DIR="$DAY"
ASM_FILE="$DIR/main.s"  # Assembly source file
OBJ_FILE="$DIR/main.o"  # Object file
BIN_FILE="$DIR/main"    # Final binary

# Check if the directory and main.s exist
if [ ! -d "$DIR" ] || [ ! -f "$ASM_FILE" ]; then
    echo "Error: Directory $DIR or file $ASM_FILE does not exist."
    exit 1
fi

# Assemble the assembly code
echo "Assembling $ASM_FILE..."
as -arch arm64 -o "$OBJ_FILE" "$ASM_FILE"
if [ $? -ne 0 ]; then
    echo "Error: Assembly failed."
    exit 1
fi

# Link the object file to create the binary
echo "Linking $OBJ_FILE to create $BIN_FILE..."
ld -o "$BIN_FILE" "$OBJ_FILE" -lSystem -syslibroot $(xcrun -sdk macosx --show-sdk-path) -e _start -arch arm64
if [ $? -ne 0 ]; then
    echo "Error: Linking failed."
    exit 1
fi

# Run the binary
echo "Running $BIN_FILE..."
"$BIN_FILE"
EXIT_CODE=$?

# Print the exit code
echo "Program exited with code $EXIT_CODE."
