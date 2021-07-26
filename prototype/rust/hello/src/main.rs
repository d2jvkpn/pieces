//
mod play;
use play::music::open as open_mp3;

//
mod archive;
use crate::archive::arch::archive_file;

//
use rand::Rng;

fn main() {
    play::music::open("a01.mp3");
    open_mp3("a02.mp3");

    play::play_mp3("Butterfly - Smile");

    archive_file("project.txt");

    let mut rng = rand::thread_rng();
    let num: i64 = rng.gen_range(1..=6);
    println!("random number: {}", num);
}
