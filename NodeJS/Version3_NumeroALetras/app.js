const express = require('express');
const { NumerosALetras } = require('numero-a-letras');

const app = express();

app.get('/', (req, res) => {

    const numero = parseInt(req.query.n);

    if (isNaN(numero)) {
        return res.send('Uso: http://localhost:3002/?n=10');
    }

    const resultado = NumerosALetras(numero);

    res.send(resultado);

});

app.listen(3002, () => {
    console.log('Servidor iniciado en http://localhost:3002');
});