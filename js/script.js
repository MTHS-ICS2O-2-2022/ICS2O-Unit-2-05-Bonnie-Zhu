// Created by: Bonnie Z
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates your pay depending on the amount of hours you work (including tax).
 */
function calculate() {
  // input
  const pay = parseFloat(document.getElementById('pay-per-hour').value)
  const hours = parseFloat(document.getElementById('hours-worked').value)

  // process
  const payment = (hours * pay) * (1.00 - 0.18)
  const Taxes = (hours * pay) - payment

  // output
  document.getElementById('payment').innerHTML = `Your pay will be: $${payment.toFixed(2)}`
  document.getElementById('Taxes').innerHTML = `The government will take: $${Taxes.toFixed(2)}`
}