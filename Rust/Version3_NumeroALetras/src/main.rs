use std::io;
use spanish_numbers::{NumberToSpanish, ScaleType};

fn main() {

    println!("Ingrese un número:");

    let mut entrada = String::new();

    io::stdin().read_line(&mut entrada).unwrap();

    let numero: u128 = entrada.trim().parse().unwrap();

    // Escala larga (España y la usada normalmente en español)
    let conversor = NumberToSpanish::new(ScaleType::Long);

    let texto = conversor.number_to_spanish(numero, " ");

    println!();

    println!("Número en español:");

    println!("{}", texto);

}