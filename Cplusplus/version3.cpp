#include <iostream>
#include <string>
#include <windows.h>
#include <wininet.h>

#pragma comment(lib, "wininet.lib")

using namespace std;

// Función auxiliar para enviar la petición SOAP y retornar la respuesta
string enviarPeticionSOAP(string numero, string idioma) {
    string soap_envelope = 
        "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n"
        "<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\r\n"
        "  <soap:Body>\r\n"
        "    <ubiNum>" + numero + "</ubiNum>\r\n"
        "    <idioma>" + idioma + "</idioma>\r\n"
        "  </soap:Body>\r\n"
        "</soap:Envelope>\r\n";

    HINTERNET hInternet = InternetOpenA("ClienteSOAP_V3", INTERNET_OPEN_TYPE_DIRECT, NULL, NULL, 0);
    if (!hInternet) return "Error: No se pudo inicializar WinINet.";

    HINTERNET hConnect = InternetConnectA(hInternet, "localhost", 4567, NULL, NULL, INTERNET_SERVICE_HTTP, 0, 0);
    if (!hConnect) {
        InternetCloseHandle(hInternet);
        return "Error: No se pudo conectar al servidor local (localhost:4567).";
    }

    HINTERNET hRequest = HttpOpenRequestA(hConnect, "POST", "/fake_soap", NULL, NULL, NULL, INTERNET_FLAG_RELOAD, 0);
    string headers = "Content-Type: text/xml; charset=utf-8\r\n";
    
    BOOL bSend = HttpSendRequestA(hRequest, headers.c_str(), headers.length(), (LPVOID)soap_envelope.c_str(), soap_envelope.length());

    string resultado = "Error: Respuesta inesperada del servidor.";
    if (bSend) {
        char buffer[2048];
        DWORD bytesRead;
        string response = "";
        
        while (InternetReadFile(hRequest, buffer, sizeof(buffer) - 1, &bytesRead) && bytesRead > 0) {
            buffer[bytesRead] = '\0';
            response += buffer;
        }

        size_t start = response.find("<NumberToWordsResult>");
        size_t end = response.find("</NumberToWordsResult>");
        if (start != string::npos && end != string::npos) {
            start += 21; 
            resultado = response.substr(start, end - start);
        }
    } else {
        resultado = "Error: El servidor de Ruby está apagado.";
    }

    InternetCloseHandle(hRequest);
    InternetCloseHandle(hConnect);
    InternetCloseHandle(hInternet);
    
    return resultado;
}

int main() {
    SetConsoleCP(CP_UTF8);
    SetConsoleOutputCP(CP_UTF8);

    string numero;
    string idioma;
    
    while (true) {
        cout << "\n=============================================" << endl;
        cout << "=== CLIENTE SOAP C++ - VERSION 3 (BUCLE)  ===" << endl;
        cout << "=============================================" << endl;
        cout << "Ingrese un numero (o escriba 'salir' para terminar): ";
        cin >> numero;
        
        if (numero == "salir" || numero == "SALIR") {
            cout << "\n¡Gracias por usar el traductor! Saliendo..." << endl;
            break;
        }
        
        cout << "Seleccione el idioma (es / en): ";
        cin >> idioma;
        
        // Enviamos los datos al servidor de Ruby usando nuestra función
        string traduccion = enviarPeticionSOAP(numero, idioma);
        
        cout << "\n----------------------------------------" << endl;
        cout << "Resultado (" << idioma << "): " << traduccion << endl;
        cout << "----------------------------------------" << endl;
        
        cout << "\nPresione cualquier tecla para realizar otra consulta..." << endl;
        system("pause > nul"); // Pausa estética limpia antes de limpiar pantalla
        system("cls");         // Limpia la pantalla para la siguiente consulta
    }
    
    return 0;
}