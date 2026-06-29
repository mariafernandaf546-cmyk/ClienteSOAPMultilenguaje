const express = require('express');
const soap = require('soap');
const translate = require('translate-google');

const app = express();

app.get('/', async (req, res) => {

    const numero = req.query.n;

    if (!numero) {
        return res.send('Uso: http://localhost:3001/?n=10');
    }

    try {

        const wsdl =
            'https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL';

        const client = await soap.createClientAsync(wsdl);

        const [resultado] =
            await client.NumberToWordsAsync({
                ubiNum: numero
            });

        const ingles = resultado.NumberToWordsResult;

        const espanol =
            await translate(ingles, {
                from: 'en',
                to: 'es'
            });

        res.send(espanol);

    } catch (error) {

        console.error(error);
        res.status(500).send(error.message);

    }

});

app.listen(3001, () => {
    console.log('Servidor iniciado en http://localhost:3001');
});