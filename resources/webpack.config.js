const ExtractTextPlugin = require('extract-text-webpack-plugin');
const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    entry: './app.js',
    output: {
        path: path.resolve(__dirname, '../assets'),
        filename: 'app.js',
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: [
                        {loader: 'css-loader', options: {importLoaders: 1}},
                        'postcss-loader'
                    ]
                })
            },
            {
                test: /\.(png|jpg|gif|svg|eot|ttf|woff|woff2)$/,
                loader: 'url-loader',
                options: {
                    limit: 10000
                }
            }
        ]
    },
    plugins: [
        new ExtractTextPlugin('styles.css'),
        new CopyWebpackPlugin([
            {
                from: './images',
                to: 'images',
            },
        ]),
    ],
};
