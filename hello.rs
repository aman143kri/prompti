use std::process::Command;

fn main() {
    let output = Command::new("touch")
        .arg("hello.txt")
        .output()
        .expect("failed to execute process");

    if output.status.success() {
        println!("hello.txt created");
    } else {
        eprintln!("Error creating file");
    }
}
