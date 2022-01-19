#!/usr/bin/env zx
let files = await fs.readdirSync('./')
let num = await question('What is the num? ')
for ( const file of files) {
    if (file.indexOf(num) !== -1) {
        const strs = file.split('.')
        const title = strs[1]
        await $`
        mkdir ${num}.${title}
        echo "" > ${num}.${title}/${num}.${title}.go
        echo "" > ${num}.${title}/README.md
        echo "" >> README.md
        mv ./${num}.${title}.go ./${num}.${title}/${num}.${title}.go
        git add *
        Git commit -m ":sparkles:feat(tree): add solution to [${num}]${title}"
        `
        fs.writeFile(`./${num}.${title}/README.md`, `#[${num}]${title}`)
        fs.writeFile(`./README.md`, `  - [x] ${num}.${title}`)
    }
}