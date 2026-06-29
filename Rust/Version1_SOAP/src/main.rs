use quick_xml::events::Event;
use quick_xml::Reader;
use reqwest::blocking::Client;
use std::io;

fn main() {

    println!("Ingrese un número:");

    let mut numero = String::new();

    io::stdin().read_line(&mut numero).unwrap();

    let numero = numero.trim();

    let soap = format!(
r#"<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:xsd="http://www.w3.org/2001/XMLSchema"
xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">

<soap:Body>

<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
<ubiNum>{}</ubiNum>
</NumberToWords>

</soap:Body>

</soap:Envelope>"#,
        numero
    );

    let client = Client::new();

    let respuesta = client
        .post("https://www.dataaccess.com/webservicesserver/NumberConversion.wso")
        .header(
            "SOAPAction",
            "\"http://www.dataaccess.com/webservicesserver/NumberToWords\"",
        )
        .header("Content-Type", "text/xml; charset=utf-8")
        .body(soap)
        .send()
        .unwrap()
        .text()
        .unwrap();

    let mut reader = Reader::from_str(&respuesta);

    reader.config_mut().trim_text(true);

    let mut texto = String::new();

    loop {

        match reader.read_event() {

            Ok(Event::Start(ref e))
                if e.name().as_ref() == b"m:NumberToWordsResult"
                || e.name().as_ref() == b"NumberToWordsResult" =>
            {
                if let Ok(Event::Text(e)) = reader.read_event() {
                    texto = e.decode().unwrap().to_string();
                    break;
                }
            }

            Ok(Event::Eof) => break,

            Err(e) => panic!("{:?}", e),

            _ => {}
        }
    }

    println!();

    println!("Número en inglés:");

    println!("{}", texto);
}