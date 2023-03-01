describe('Home Screen', () => {
  it('should login successfully', () => {
    cy.visit('/'); // assumes the home screen URL is the root URL of the app

    cy.get('[data-test=username-input]').type('testuser');
    cy.get('[data-test=password-input]').type('testpassword');
    cy.get('[data-test=login-button]').click();

    cy.url().should('include', 'screen-search');
  });
});
