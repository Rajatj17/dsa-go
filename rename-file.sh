#!/bin/bash

# Script to convert numbered prefixes to alphabetical prefixes
# Example: "1_basic" -> "a_basic", "2_advanced" -> "b_advanced"
# Usage: ./number_to_alphabet.sh [files|dirs|both]
# Default: both

# Function to convert number to alphabet
number_to_alphabet() {
    local num=$1
    if [ $num -le 26 ]; then
        # Single letter for numbers 1-26
        printf "\\$(printf '%03o' $((96 + num)))"
    else
        # Double letters for numbers > 26 (aa, ab, ac, etc.)
        local first=$(( (num - 1) / 26 ))
        local second=$(( (num - 1) % 26 + 1 ))
        printf "\\$(printf '%03o' $((96 + first)))\\$(printf '%03o' $((96 + second)))"
    fi
}

# What to rename
TARGET="both"
if [ "$1" = "files" ] || [ "$1" = "dirs" ] || [ "$1" = "both" ]; then
    TARGET="$1"
elif [ ! -z "$1" ]; then
    echo "Usage: $0 [files|dirs|both]"
    echo "  files: Rename only files"
    echo "  dirs: Rename only directories"
    echo "  both: Rename both files and directories (default)"
    exit 1
fi

echo "Converting numbered prefixes to alphabetical prefixes..."
echo "Target: $TARGET"
echo

renamed_count=0

# Function to rename items
rename_item() {
    local item_path="$1"
    local item_type="$2"
    
    local dir=$(dirname "$item_path")
    local name=$(basename "$item_path")
    
    # Check if name starts with a number followed by underscore, dot, or dash
    if [[ $name =~ ^([0-9]+)([_.+-])(.+)$ ]]; then
        local number="${BASH_REMATCH[1]}"
        local separator="${BASH_REMATCH[2]}"
        local rest="${BASH_REMATCH[3]}"
        
        # Convert number to alphabet
        local alphabet=$(number_to_alphabet $number)
        
        # Create new name
        local new_name="${alphabet}${separator}${rest}"
        local new_path="$dir/$new_name"
        
        # Check if target already exists
        if [ -e "$new_path" ]; then
            echo "⚠️  Skipping $item_type: $item_path"
            echo "   Target already exists: $new_path"
            echo
        else
            # Rename the item
            mv "$item_path" "$new_path"
            echo "✅ Renamed $item_type: $item_path"
            echo "   To: $new_path"
            echo
            ((renamed_count++))
        fi
    fi
}

# Rename directories first (deepest first to avoid path issues)
if [ "$TARGET" = "dirs" ] || [ "$TARGET" = "both" ]; then
    echo "🔄 Renaming directories..."
    find . -depth -type d | while read -r dir; do
        rename_item "$dir" "directory"
    done
fi

# Rename files
if [ "$TARGET" = "files" ] || [ "$TARGET" = "both" ]; then
    echo "🔄 Renaming files..."
    find . -type f | while read -r file; do
        rename_item "$file" "file"
    done
fi

echo "Conversion complete!"
echo "Items renamed: $renamed_count"
echo
echo "📝 Examples of conversions:"
echo "   1_basic → a_basic"
echo "   2_advanced → b_advanced"
echo "   10_example → j_example"
echo "   27_test → aa_test"
echo "   52_large → az_large"