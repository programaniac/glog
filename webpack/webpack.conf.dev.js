const path = require("path");

const getHTMLFiles = require("./utils/get-html-files");
const parseFiles = require("./utils/parse-files");
const getJavascriptFiles = require("./utils/get-javascript-files");
const webpack = require("webpack")
const autoprefixer = require('autoprefixer');

const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const extractLESS = new ExtractTextPlugin('css/[name].css');

const config = {
	entry: {},
	output: {
		path: path.resolve(__dirname, "../out/public"),
		filename: 'js/[name].js',
		publicPath: '/',
		chunkFilename: 'js/[id].js'
	},
	mode: "development",
	devtool: 'source-map',
	watch: true,
	resolve: {
		extensions: [".js", ".jsx", ".ts", ".tsx", ".css", ".less"]
	},
	module: {
		rules: [{
				test: /\.js?x$/,
				exclude: /(node_modules|bower_components)/,
				use: {
					loader: 'babel-loader',
					options: {}
				}
			},
			{
				test: /\.tsx?$/,
				loader: "ts-loader"
			},
			{
				test: /\.html$/,
				use: [{
					loader: 'html-loader',
					options: {
						interpolate: true
					}
				}]
			},
			{
				test: /\.less$/,
				use: extractLESS.extract(['css-loader',
					{
						loader: 'postcss-loader',
						options: {
							plugins: () => autoprefixer({
								browsers: ['last 3 versions', '> 0.1%']
							})
						}
					}, 'less-loader'
				])
			},
			{
				test: /\.css$/,
				use: ['style-loader', 'css-loader'],
			},
			{
				test: /\.(gif|png|jpe?g|svg)$/i,
				use: [{
					loader: 'file-loader',
					options: {
						name: "img/[name].[ext]"
					}
				}],
			}
		]
	},
	plugins: [
		new CleanWebpackPlugin([
			"public",
			"views"
		], {
			root: path.resolve(__dirname, "../out"),
			dry: false,
			verbose: true
		}),
		extractLESS,
		new webpack.ProvidePlugin({ $: 'jquery', jQuery: 'jquery' }),
		new webpack.DefinePlugin({
			'process.env': {
				'NODE_ENV': JSON.stringify(process.env.NODE_ENV)
			},
		})
	]
}

const clientFolder = path.resolve(__dirname, "../client/")

const htmlFiles = getHTMLFiles(clientFolder);
const parsedHTMLFiles = parseFiles(htmlFiles);

const jsFiles = getJavascriptFiles(parsedHTMLFiles);
const parsedJSFiles = parseFiles(jsFiles)

parsedJSFiles.forEach(jsFile => {
	config.entry[jsFile.name] = jsFile.dir + "/" + jsFile.base;
});

parsedHTMLFiles.forEach(htmlFile => {
	config.plugins.push(new HtmlWebpackPlugin({
		template: htmlFile.dir + "/" + htmlFile.base,
		filename: path.resolve(__dirname, "../out/views/" + htmlFile.name + ".html"),
		chunks: [htmlFile.name],
		inject: true,
		minify: false
	}))
})


module.exports = config;