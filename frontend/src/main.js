import 'core-js/stable';
const runtime = require('@wailsapp/runtime');

function start() {
	const formCalculo = document.querySelector("#formCalculo");
	const peso = document.querySelector("#peso");
	const copo = document.querySelector("#copo");

	var app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	formCalculo.addEventListener('submit', (e) => {
		e.preventDefault();
		console.log("Peso: " + peso.value + " ; Copo: " + copo.value);
	});
};

runtime.Init(start);
