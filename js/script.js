// Created by: Bonnie Z
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates your pay depending on the amount of hours you work (including tax).
 */
function calculate() {
  // input
  const pay = parseInt(document.getElementById('pay-per-hour').value)
  const hours = parseInt(document.getElementById('hours-worked').value)

  // process
  const payment = (hours * pay) * (1.00 - 0.18)
  const taxes = (hours * pay) * 0.18

  // output
  document.getElementById('payment').innerHTML = `Your pay will be: $ ${payment.toFixed(2)}`
  document.getElementById('tax').innerHTML = `the government will take: $ ${taxes.toFixed(2)}`
}