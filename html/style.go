package wikiwikihtml


	const style string =
`html, body {
	height: 100%;
}

body {
	font-family: serif;
	display: flex;
	margin: auto;
	justify-content: center;
}
body {
	color: #333333;
	background-color:white;
}

main {
	margin: 75pt 11.25pt;
	max-width: 502.5pt;

	@media screen and (min-width: 450pt){
		padding: 0 22.5pt;
	}
}

h1,h2,h3,h4,h5,h6 {
	color: #111111;
	font-weight: 700;
}
h1 {
	text-align:center;
}
h2 {
	border-bottom:1px solid #d0d0d0;
	margin:21pt 0 0 0;
	padding-bottom:3pt;
}
h3 {
	text-transform:uppercase;
}

blockquote {
	background-color: rgba(0,0,0, 0.07);
	color:            rgba(0,0,0, 0.8);

	margin:  0.25em 0.25em;
	padding: 1.25em 1.25em;

	border-left:  0.25em solid rgba(0,0,0, 0.1);
	border-right: 0.25em solid rgba(0,0,0, 0.1);
}

p {
	font-size: 12pt;
	line-height: 1.625;
	word-wrap: break-word;
	overflow-wrap: break-word;
	hyphens: auto;

	@media screen and (min-width: 450pt){
		font-size: 13.5pt;
		line-height: 1.667;
	}
}
p, li, blockquote {
	text-align: justify;
}

img[align="right"] {
	margin-left:1.5em;
}

body {
	counter-reset: citation-number-counter;
}
sup.citation a:before {
	counter-increment: citation-number-counter +1;
	content: "[" counter(citation-number-counter) "]";
}
`+
`
a.wiki-link::before{
	content:attr(href);
}
`
