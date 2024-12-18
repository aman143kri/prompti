import java.io.*;

public class RunCommand {
    public static void main(String[] args) {
        try {
            Process process = Runtime.getRuntime().exec("touch hello.txt");
            process.waitFor();
            System.out.println("hello.txt created");
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}
