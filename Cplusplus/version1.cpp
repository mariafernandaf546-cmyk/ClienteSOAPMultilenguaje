#include <iostream>
#include <string>
#include <windows.h>
#include <wininet.h>

#pragma comment(lib, "wininet.lib")

using namespace std;

int main() {
    SetConsoleCP(CP_UTF8);
    SetConsoleOutputCP(CP_UTF8);

    string numero;
    
    cout << "=== CLIENTE SOAP C++ - VERSION 1 ===" << endl;
    cout << "Ingrese un numero: ";
    cin >> numero;

    // En la Versión 1 solo se envía el número, no existe la etiqueta <idioma>
    string soap_envelope = 
        "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n"
        "<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\r\n"
        "  <soap:Body>\r\n"
        "    <ubiNum>" + numero + "</ubiNum>\r\n"
        "  </soap:Body>\r\n"
        "</soap:Envelope>\r\n";

    HINTERNET hInternet = InternetOpenA("ClienteSOAP_V1", INTERNET_OPEN_TYPE_DIRECT, NULL, NULL, 0);
    HINTERNET hConnect = InternetConnectA(hInternet, "localhost", 4567, NULL, NULL, INTERNET_SERVICE_HTTP, 0, 0);
    HINTERNET hRequest = HttpOpenRequestA(hConnect, "POST", "/fake_soap", NULL, NULL, NULL, INTERNET_FLAG_RELOAD, 0);

    string headers = "Content-Type: text/xml; charset=utf-8\r\n";
    BOOL bSend = HttpSendRequestA(hRequest, headers.c_str(), headers.length(), (LPVOID)soap_envelope.c_str(), soap_envelope.length());

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
            string resultado = response.substr(start, end - start);
            
            cout << "\n========================================" << endl;
            cout << "Resultado: " << resultado << endl;
            cout << "========================================" << endl;
        } else {
            cout << "\nError: Respuesta inesperada del servidor." << endl;
        }
    } else {
        cout << "\nError: No se pudo conectar con el servidor." << endl;
    }

    InternetCloseHandle(hRequest);
    InternetCloseHandle(hConnect);
    InternetCloseHandle(hInternet);
    
    return 0;
}