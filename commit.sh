#! /bin/sh

echo "Enter message"
read -r msg

git add .
git commit -m "$msg"
git push