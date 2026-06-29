const express = require('express');
const soap = require('soap');

const app = express();

app.get('/', async (req, res) => {

    const numero = req.query.n;

    if (!numero) {
        return res.send('Uso: http://localhost:3000/?n=10');
    }

    const wsdl = 'https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL';

    try {

        const client = await soap.createClientAsync(wsdl);

        const [resultado] = await client.NumberToWordsAsync({
            ubiNum: numero
        });

        res.send(resultado.NumberToWordsResult);

    } catch (error) {

        console.error(error);
        res.status(500).send(error.message);

    }

});

app.listen(3000, () => {
    console.log('Servidor iniciado en http://localhost:3000');
});