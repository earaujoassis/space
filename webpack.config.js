const path = require('path')
const webpack = require('webpack')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const TerserPlugin = require('terser-webpack-plugin')
const commitHash = process.env.COMMIT_HASH || 'unknown-version'

module.exports = {
    entry: {
        himalia: './web/himalia/index.js',
        ganymede: './web/ganymede/index.js'
    },
    output: {
        path: path.resolve(__dirname, './web/public/'),
        publicPath: '/public/',
        filename: 'js/[name].min.js'
    },
    resolve: {
        alias: {
            'public-css': path.resolve(__dirname, 'web/public/css'),
            '@core': path.resolve(__dirname, 'web/core'),
            '@app': path.resolve(__dirname, 'web/himalia/app'),
            '@actions': path.resolve(__dirname, 'web/himalia/actions'),
            '@components': path.resolve(__dirname, 'web/himalia/components'),
            '@containers': path.resolve(__dirname, 'web/himalia/containers'),
            '@stores': path.resolve(__dirname, 'web/himalia/stores'),
            '@ui': path.resolve(__dirname, 'web/himalia/ui'),
            '@utils': path.resolve(__dirname, 'web/himalia/utils'),
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
                    }
                ]
            }
        ]
    }
}
