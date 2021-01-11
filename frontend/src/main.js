import 'core-js/stable';
const runtime = require('@wailsapp/runtime');

function start() {
	var app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	app.innerHTML = `
	<header id="header">
						<img src="garrafinha.png" /> 
						<span>Hidrate-se</span>
					</header>
					<main id="content">
						<p id="informativo">Antes de iniciar o app, precisamos saber algumas informações sobre você.</p>
						<form id="formCalculo">
							<div class="form-input">
								<label for="peso">Peso (kg): </label>
								<input type="number" id="peso" name="peso" required> 
							</div>
							<div class="form-input">
								<label for="copo">Copo (ml): </label>
								<input type="number" id="copo" name="copo" required> 
							</div>
							<button>Começar</button>
						</form>
						<div id="resposta">
						</div>
					</main>
`;

	const formCalculo = document.querySelector("#formCalculo");
	const peso = document.querySelector("#peso");
	const copo = document.querySelector("#copo");
	const resposta = document.querySelector("#resposta");


	formCalculo.addEventListener('submit', (e) => {
		e.preventDefault();
		console.log("Peso: " + peso.value + " ; Copo: " + copo.value);
		window.backend.Contador.Calcular(parseFloat(peso.value), parseFloat(copo.value))
			.then(res => {
				document.querySelector("#informativo").style.display = 'none';
				formCalculo.style.display = 'none';
				resposta.innerHTML = res;
			})
	});
};

runtime.Init(start);
