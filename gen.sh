mkdir $1.$2
echo "" > $1.$2/$1.$2.go
echo "#[$1]$2" > $1.$2/README.md
echo "  - [x] $1.$2" >> README.md
mv ./$1.$2.go ./$1.$2/$1.$2.go
git add *
git commit -m ":sparkles:feat($3): add solution to [$1]$2"