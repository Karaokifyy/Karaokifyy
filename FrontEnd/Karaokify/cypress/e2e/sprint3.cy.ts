describe('Test 1', () => {
    it('Get Songs from playlist', () => {
      cy.visit('/songs')
      //cy.contains('Welcome')
    })
  })




describe('Test 2', () => {
    it('Full Login Complete', () => {
        cy.visit('/')
        cy.get('[name="username"]').type('skandavskumaran@gmail.com');
        cy.get('[name="password"]').type('ayyampet');
        //cy.contains('Sign In').click();
        //cy.contains('LOG IN').click()
        // cy.url()
        //.should('includes','/screen-search')

    })
  })


  describe('Test 3', () => {
    it('Full Process to get Karaoke from spotify playlist', () => {
      cy.visit('/lyrics')
      //cy.get('button').click()
    })
  })  
