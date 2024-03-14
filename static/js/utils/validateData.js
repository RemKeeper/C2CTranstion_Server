function validatePasswordStrength(password) {
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/;
    return passwordRegex.test(password) && password.length >= 16;
}

