const HTML2PDF = require('./html2pdf.js');

var generate = async () => {
    await HTML2PDF._initializePlugins();
    await HTML2PDF.pdf('./knuth.pug');
};

generate();
