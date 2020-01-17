const { masterToPDF } = require('relaxedjs');
const puppeteer = require('puppeteer');
const plugins = require('relaxedjs/src/plugins');

class HTML2PDF {
    constructor(){
        this.puppeteerConfig = {
            headless: true,
            args: [
                '--no-sandbox',
                '--disable-translate',
                '--disable-extensions',
                '--disable-sync'
            ],
        };

        this.relaxedGlobals = {
            busy: false,
            config: {},
            configPlugins: [],
        };
    }

    async _initializePlugins() {
        for (let [i, plugin] of plugins.builtinDefaultPlugins.entries()) {
            plugins.builtinDefaultPlugins[i] = await plugin.constructor();
        }
        await plugins.updateRegisteredPlugins(this.relaxedGlobals, '/');

        let chrome = await puppeteer.launch(this.puppeteerConfig);
        this.relaxedGlobals.puppeteerPage = await chrome.newPage();
    }

    async pdf(template_pug, json_data){
        //await this._initPDFReader();
        
        await masterToPDF(template_pug,
            this.relaxedGlobals,
            __dirname+'/knuth_temp.htm',
            __dirname+'/knuth.pdf',
            json_data,
        );
    }
}

module.exports = new HTML2PDF();