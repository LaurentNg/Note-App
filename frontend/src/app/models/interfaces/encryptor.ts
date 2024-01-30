import { EncryptorDto } from "../dtos/encryptorDto";

export class Encryptor {
    text: string;
    passphrase: string;
  
    constructor(text: string, passphrase: string) {
        this.text = text;
        this.passphrase = passphrase;
    }
  
    toDto(): EncryptorDto {
      return new EncryptorDto(this.text, this.passphrase)
    }

    obj(): object {
        return {
            text: this.text,
            passphrase: this.passphrase
        }
    }
}