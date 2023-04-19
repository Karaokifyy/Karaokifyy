describe('Test 1', () => {
    it('Navigate Pages Without Errors', () => {
      cy.visit('/songs')
      //cy.contains('Welcome')
    })
  })




describe('Test 2', () => {
    it('OAuth Login with 3rd party', () => {
        cy.visit('')
        cy.get('button').contains('Sign In').click();

        cy.origin('https://accounts.spotify.com/en/login', () => {
              cy.get('#login-username').type("skandavskumaran@gmail.com")
              cy.get('#login-password').type('ayyampet')
              //cy.contains('button[value=default]', 'Continue').click()
              //cy.get('#login-button').click();
              //cy.visit('/screen-search')
            }
          );
        //cy.get('[name="username"]').type('skandavskumaran@gmail.com');
        //cy.get('[name="password"]').type('ayyampet');
        //cy.contains('Sign In').click();
        //cy.contains('LOG IN').click()
        // cy.url()
        //.should('includes','/screen-search')

    })
  })


  describe('Test 3', () => {
    it('Use login to select playlist and song and navigate to karaoke Center', () => {
        cy.visit('/lyrics')

    })
  })
