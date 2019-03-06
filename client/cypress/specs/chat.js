// https://docs.cypress.io/api/introduction/api.html

describe('Chat test', () => {
  it('Authenticate user', () => {
    cy.visit('localhost:8080')

    cy.get('#modal_screen #inputs > #username').type('user')
    cy.get('#modal_screen #inputs > #set_username').click()
  })

  it('Send message', () => {
    cy.get('#chat #inputs > #message_box').type('test message')
    cy.get('#chat #inputs > #send_message').click()
  })

  it('Find message', () => {
    cy.get('#chat #messages').contains('test message')
  })
})
