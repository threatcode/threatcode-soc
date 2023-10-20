

require('../test_common.js');
require('./login.js');

let comp;

beforeEach(() => {
  comp = getComponent("login");
  resetPapi();
});

test('shouldSubmitTotp', () => {
  var submitted = false;
  const getElementByIdMock = global.document.getElementById = jest.fn().mockReturnValueOnce({value:""}).mockReturnValueOnce({submit:function(){ submitted = true}});
  comp.submitTotp('123');
  expect(comp.form.totpCode).toBe('123');
  expect(getElementByIdMock).toHaveBeenCalledWith('totp_code');
  expect(getElementByIdMock).toHaveBeenCalledWith('loginForm');
});