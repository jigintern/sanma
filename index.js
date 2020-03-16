const rp = require('request-promise');
const cheerio = require('cheerio');

const options = {
    transform: (body) => {
        return cheerio.load(body);
    }
};

rp.get('https://fukuno.jig.jp/', options)
    .then(($) => {
        $('[itemtype="http://schema.org/Article"]').each((id, elem) => {
            console.log($(elem).children('[itemprop=""]').text());
        });
    }).catch((error) => {
    console.log('error\n' + error);
});