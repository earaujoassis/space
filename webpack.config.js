const path = require('path')
const webpack = require('webpack')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const TerserPlugin = require('terser-webpack-plugin')
const commitHash = process.env.COMMIT_HASH || 'unknown-version'

module.exports = {
    entry: './web/himalia/index.js',
    output: {
        path: path.resolve(__dirname, './web/public/'),
        publicPath: '/public/',
        filename: 'js/himalia.min.js'
    },
    resolve: {
        alias: {
            'public-css': path.resolve(__dirname, 'web/public/css'),
            '@app': path.resolve(__dirname, 'web/himalia/app'),
            '@actions': path.resolve(__dirname, 'web/himalia/actions'),
            '@components': path.resolve(__dirname, 'web/himalia/components'),
            '@containers': path.resolve(__dirname, 'web/himalia/containers'),
            '@stores': path.resolve(__dirname, 'web/himalia/stores')
        }
    },
    devtool: 'cheap-module-source-map',
    plugins: [
        new MiniCssExtractPlugin({
            filename: 'css/himalia.css',
            chunkFilename: '[id].css'
        }),
        new webpack.DefinePlugin({
            __COMMIT_HASH__: JSON.stringify(commitHash)
        })
    ],
    optimization: {
        minimizer: [new TerserPlugin({ extractComments: false })]
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader'
                }
            },
            {
                test: /\.css$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    {
                        loader: 'css-loader',
                        options: {
                            url: false
                        }
                    },
                    'resolve-url-loader'
                ]
            }
        ]
    }
}
