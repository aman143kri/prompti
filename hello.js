const { exec } = require('child_process');

exec('touch hello.txt', (error, stdout, stderr) => {
    if (error) {
        console.error(`exec error: ${error}`);
        return;
    }
    console.log('hello.txt created');
});
