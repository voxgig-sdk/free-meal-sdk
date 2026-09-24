"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FreeMealError = void 0;
class FreeMealError extends Error {
    isFreeMealError = true;
    sdk = 'FreeMeal';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.FreeMealError = FreeMealError;
//# sourceMappingURL=FreeMealError.js.map