const express = require('express');
const converter = require('number-to-words');

const app = express();

app.get('/', (req, res) => {

    const numero = parseInt(req.query.n);

    if (isNaN(numero)) {
        return res.send('Ingrese un número. Ejemplo: ?n=10');
    }

    const texto = converter.toWords(numero);

    res.send(texto);

});

app.listen(3000, () => {
    console.log('Servidor ejecutándose en http://localhost:3000');
});