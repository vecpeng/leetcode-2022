#!/usr/bin/env zx
let files = await fs.readdirSync('./')
let num = await question('What is the num? ')
for ( const file of files) {
    if (file.indexOf(num) !== -1) {
        const strs = file.split('.')
        const title = strs[1]
        await $`
        mkdir ${num}.${title}
        echo "" > ${num}.${title}/${num}.${title}.cpp
        echo "" > ${num}.${title}/README.md
        echo "" >> README.md
        mv ./${num}.${title}.cpp ./${num}.${title}/${num}.${title}.cpp
        git add ./
        `
        fs.writeFileSync(`./${num}.${title}/README.md`, `#[${num}]${title}`)
        fs.appendFileSync(`./README.md`, `\t- [x] ${num}.${title}`)

        await $`git commit -m ${":sparkles:feat(tree): add solution to [" + num + "]" + title}`
    }
}