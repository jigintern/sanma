const rp = require('request-promise');
const cheerio = require('cheerio');
const mysql = require('mysql');

const connection = mysql.createConnection({
    host: 'test.cnddxi4etcjb.us-east-2.rds.amazonaws.com',
    user: 'admin',
    password: 'password',
    port: 3306,
    database: 'test'
});

const connection = mysql.createConnection({
    host: 'sanma-db',
    user: 'root',
    password: 'password',
    port: 3306,
    database: 'articles'
});

let targets = ['https://fukuno.jig.jp/'];

const options = {
    transform: (body) => {
        return cheerio.load(body);
    }
};

targets.forEach((target) => {
    rp.get(target, options)
        .then(($) => {
            const title = [];
            const urls = [];
            $('[itemtype="http://schema.org/Article"]').each((id, elem) => {
                title.push('\'' + $(elem).find('[itemprop="headline"]').text() + '\'');
                urls.push('\'' + $(elem).find('[itemprop="url"]').attr('href') + '\'');
                //author.push($(elem).find('[itemprop="author"]').text())
                //status: {draft: 0, publish: 1}
                //type: {origin: 0, crawled: 1}
            });
            connection.connect();
            for (var count = 0; count < title.length; count++) {
                connection.query('INSERT INTO articles (article_status, article_type, article_url, title, updated_at) VALUES (1, 1, ' + urls[count] + ', ' + title[count] + ', CURRENT_TIMESTAMP)',
                    (err, results, field) => {
                        if (err) throw err;
                    });
            }
        }).catch((error) => {
        console.log('error\n' + error);
    }).finally(() => {
        connection.end();
    });
});
