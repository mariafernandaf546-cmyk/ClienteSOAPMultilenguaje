import java.io.*;
import java.net.*;
import java.util.Scanner;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;

import org.jsoup.Jsoup;
import org.w3c.dom.Document;
import org.w3c.dom.NodeList;
import org.xml.sax.InputSource;

public class ClienteV2 {

    public static void main(String[] args) {

        try {

            Scanner sc = new Scanner(System.in);

            System.out.print("Ingrese un número: ");
            int numero = sc.nextInt();

            String ingles = convertirNumero(numero);

            System.out.println("\nNúmero en inglés:");
            System.out.println(ingles);

            String espanol = traducir(ingles);

            System.out.println("\nNúmero en español:");
            System.out.println(espanol);

        } catch (Exception e) {
            e.printStackTrace();
        }

    }

    public static String convertirNumero(int numero) throws Exception {

        String soap =
                "<?xml version=\"1.0\" encoding=\"utf-8\"?>" +
                "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" " +
                "xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" " +
                "xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">" +
                "<soap:Body>" +
                "<NumberToWords xmlns=\"http://www.dataaccess.com/webservicesserver/\">" +
                "<ubiNum>" + numero + "</ubiNum>" +
                "</NumberToWords>" +
                "</soap:Body>" +
                "</soap:Envelope>";

        URL url = new URL("https://www.dataaccess.com/webservicesserver/NumberConversion.wso");

        HttpURLConnection con = (HttpURLConnection) url.openConnection();

        con.setRequestMethod("POST");

        con.setRequestProperty("Content-Type", "text/xml; charset=utf-8");

        con.setRequestProperty(
                "SOAPAction",
                "\"http://www.dataaccess.com/webservicesserver/NumberToWords\""
        );

        con.setDoOutput(true);

        OutputStream os = con.getOutputStream();
        os.write(soap.getBytes("UTF-8"));
        os.close();

        BufferedReader br = new BufferedReader(
                new InputStreamReader(con.getInputStream()));

        String linea;

        StringBuilder xml = new StringBuilder();

        while ((linea = br.readLine()) != null) {
            xml.append(linea);
        }

        br.close();

        DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();

        factory.setNamespaceAware(true);

        DocumentBuilder builder = factory.newDocumentBuilder();

        Document doc = builder.parse(
                new InputSource(
                        new StringReader(xml.toString())
                )
        );

        NodeList lista = doc.getElementsByTagNameNS("*", "NumberToWordsResult");

        return lista.item(0).getTextContent().trim();

    }

    public static String traducir(String texto) throws Exception {

        String url =
                "https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=es&dt=t&q="
                        + URLEncoder.encode(texto, "UTF-8");

        org.jsoup.nodes.Document pagina =
                Jsoup.connect(url)
                        .ignoreContentType(true)
                        .userAgent("Mozilla/5.0")
                        .timeout(15000)
                        .get();

        String json = pagina.body().text();

        System.out.println("\nRespuesta de Google:");
        System.out.println(json);

        int inicio = json.indexOf("\"") + 1;
        int fin = json.indexOf("\"", inicio);

        if (inicio <= 0 || fin <= inicio)
            return "No fue posible traducir.";

        return json.substring(inicio, fin);

    }

}