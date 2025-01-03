
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});
function grade() {
    rl.on('line', (input) => {
        console.log('input :>> ', input);
    })
}