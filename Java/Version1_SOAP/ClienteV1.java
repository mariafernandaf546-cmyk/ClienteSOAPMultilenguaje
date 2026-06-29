import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class ClienteV1 {

    public static void main(String[] args) {
        Scanner teclado = new Scanner(System.in);
        
        System.out.println("=== CLIENTE SOAP JAVA - VERSION 1 (DIRECTO A INTERNET) ===");
        System.out.print("Ingrese un numero: ");
        String numero = teclado.nextLine();

        String soapEnvelope = 
            "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n" +
            "<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\r\n" +
            "  <soap:Body>\r\n" +
            "    <NumberToWords xmlns=\"http://www.dataaccess.com/webservicesserver/\">\r\n" +
            "      <ubiNum>" + numero + "</ubiNum>\r\n" +
            "    </NumberToWords>\r\n" +
            "  </soap:Body>\r\n" +
            "</soap:Envelope>\r\n";

        try {
            URL url = new URL("https://www.dataaccess.com/webservicesserver/NumberConversion.wso");
            HttpURLConnection conexion = (HttpURLConnection) url.openConnection();
            
            conexion.setRequestMethod("POST");
            conexion.setRequestProperty("Content-Type", "text/xml; charset=utf-8");
            conexion.setDoOutput(true);

            try (DataOutputStream wr = new DataOutputStream(conexion.getOutputStream())) {
                wr.writeBytes(soapEnvelope);
                wr.flush();
            }

            int responseCode = conexion.getResponseCode();
            if (responseCode == 200) {
                BufferedReader in = new BufferedReader(new InputStreamReader(conexion.getInputStream()));
                String inputLine;
                StringBuilder respuestaCompleta = new StringBuilder();

                while ((inputLine = in.readLine()) != null) {
                    respuestaCompleta.append(inputLine);
                }
                in.close();

                String xmlResponse = respuestaCompleta.toString();

                // EXTRACCIÓN ULTRA-ROBUSTA: Usamos Expresiones Regulares para capturar 
                // lo que esté dentro de NumberToWordsResult, ignore o no los namespaces intermedios.
                Pattern pattern = Pattern.compile("<[^>]*NumberToWordsResult[^>]*>(.*?)</[^>]*NumberToWordsResult>");
                Matcher matcher = pattern.matcher(xmlResponse);

                if (matcher.find()) {
                    String resultado = matcher.group(1);
                    
                    System.out.println("\n========================================");
                    System.out.println("Resultado (Internet): " + resultado.trim());
                    System.out.println("========================================");
                } else {
                    System.out.println("\nError: No se encontró la etiqueta de resultado en el XML.");
                }
            } else {
                System.out.println("\nError: El servidor externo respondió con código " + responseCode);
            }

        } catch (Exception e) {
            System.out.println("\nError de red: " + e.getMessage());
        }
        
        teclado.close();
    }
}