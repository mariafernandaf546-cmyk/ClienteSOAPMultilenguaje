import java.util.Scanner;

import com.ibm.icu.text.RuleBasedNumberFormat;

public class ClienteV3 {

    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);

        System.out.print("Ingrese un número: ");

        long numero = sc.nextLong();

        RuleBasedNumberFormat formato =
                new RuleBasedNumberFormat(
                        new java.util.Locale("es", "ES"),
                        RuleBasedNumberFormat.SPELLOUT);

        String resultado = formato.format(numero);

        System.out.println();
        System.out.println("Número en español:");
        System.out.println(resultado);

        sc.close();
    }
}