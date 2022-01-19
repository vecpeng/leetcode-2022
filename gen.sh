files=$(ls ./)
title=""
for file in $files
do
#   if [ $(expr index ${file} $1) != 0 ]
#   then
#     title=${file: index : }
#     echo title
    index=$(expr index ${file} $1)
    echo $index
    if test -f $file && [ $index != 0 ]
    then
        echo $file
    fi
done
mkdir $1.$2
echo "" > $1.$2/$1.$2.go
echo "#[$1]$2" > $1.$2/README.md
echo "  - [x] $1.$2" >> README.md
mv ./$1.$2.go ./$1.$2/$1.$2.go
git add *
git commit -m ":sparkles:feat($3): add solution to [$1]$2"