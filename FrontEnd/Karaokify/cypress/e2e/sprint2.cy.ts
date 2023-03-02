describe('Starter Test', () => {
    it('Visits the initial project page', () => {
      cy.visit('/')
      cy.contains('Welcome')
    })
  })


describe('Test 2', () => {
    it('Visits the search bar page', () => {
      cy.visit('/screen-search')
    })
  })

describe('Test 3', () => {
    it('Get Username and Password', () => {
      cy.visit('/')
      cy.get('[name="username"]').type('test');
      cy.get('[name="password"]').type('test');
      //cy.get('button').click()
    })
  })