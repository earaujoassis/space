const path = require('path')
const webpack = require('webpack')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const TerserPlugin = require('terser-webpack-plugin')
const RemoveEmptyScriptsPlugin = require('webpack-remove-empty-scripts')
const commitHash = process.env.COMMIT_HASH || 'unknown-version'
const environment = process.env.NODE_ENV || 'development'
const isProduction = environment === 'production'

module.exports = {
    entry: {
        amalthea: ['./web/amalthea/index.js', './web/amalthea/styles/amalthea.scss'],
        callisto: ['./web/callisto/index.js', './web/callisto/styles/callisto.scss'],
        ganymede: ['./web/ganymede/index.js', './web/ganymede/styles/ganymede.scss'],
        himalia: './web/himalia/index.js',
        io: ['./web/io/index.js', './web/io/styles/io.scss'],
        core: './web/core/styles/core.scss',
        errors: './web/core/styles/errors.scss'
    },
    output: {
        path: path.resolve(__dirname, './web/public/'),
        publicPath: '/public/',
        filename: 'js/[name].min.js',
        clean: false
    },
    resolve: {
        alias: {
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
        new RemoveEmptyScriptsPlugin(),
        new MiniCssExtractPlugin({
            filename: 'css/[name].css',
            chunkFilename: 'css/[name].css'
        }),
        new webpack.DefinePlugin({
            __COMMIT_HASH__: JSON.stringify(commitHash)
        })
    ],
    optimization: {
        minimizer: [new TerserPlugin({ extractComments: false })]
    },
    watchOptions: {
        ignored: /node_modules/,
        aggregateTimeout: 200,
        poll: false
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
            },
            {
                test: /\.scss$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    {
                        loader: 'sass-loader',
                        options: {
                            sassOptions: {
                                includePaths: [
                                    'node_modules/foundation-sites/scss',
                                    'node_modules/normalize.scss',
                                    'web/core/styles',
                                    'web'
                                ],
                                outputStyle: isProduction ? 'compressed' : 'expanded',
                                quietDeps: true,
                                silenceDeprecations: ['legacy-js-api', 'import', 'global-builtin', 'mixed-decls']
                            }
                        }
                    }
                ]
            }
        ]
    }
}
