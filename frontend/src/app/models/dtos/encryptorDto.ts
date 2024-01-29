import { Encryptor } from "../interfaces/encryptor";

export class EncryptorDto {
    text: string;
    passphrase: string;
  
    constructor(text: string, passphrase: string) {
        this.text = text;
        this.passphrase = passphrase;
    }
  
    toInterface(): Encryptor {
      return new Encryptor(this.text, this.passphrase)
    }

    obj(): object {
        return {
            text: this.text,
            passphrase: this.passphrase
        }
    }
}