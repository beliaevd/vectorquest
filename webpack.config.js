module.exports = {
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            }
        ]
    },
    entry: {
        shared : "./web/static/shared/js/*.js",
    },
    output: {
        path: __dirname + '/web/static/',
        publicPath: '',
        filename: '/build/js/[name].js'
    },
    devServer: {
        contentBase: './distgetuk'
    }
};