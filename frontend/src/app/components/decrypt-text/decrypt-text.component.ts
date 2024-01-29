import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Clipboard } from '@angular/cdk/clipboard';
import { EncryptorHttpService } from 'src/app/services/http/encryptor-http.service';
import { Encryptor } from 'src/app/models/interfaces/encryptor';


@Component({
  selector: 'app-decrypt-text',
  templateUrl: './decrypt-text.component.html',
  styleUrl: './decrypt-text.component.scss'
})
export class DecryptTextComponent implements OnInit {
  notesToDecrypt = new FormGroup({
    text: new FormControl<string>('', [Validators.required]),
    passphrase: new FormControl<string>('', [Validators.required])
  });
  decryptedText = ""
  constructor(private clipboard: Clipboard, private encryptor: EncryptorHttpService) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.notesToDecrypt.valid) {
      const text = this.notesToDecrypt.get('text')?.value
      const passphrase = this.notesToDecrypt.get('passphrase')?.value
      const encryptorObj = new Encryptor(text as string, passphrase as string)
      this.encryptor.decrypt(encryptorObj).subscribe((res) => {
        if ('text' in res) this.decryptedText = res.text
      })
    }
    else console.log('not valid')
  }

  copyDecryptedText() {
    const pending = this.clipboard.beginCopy(this.decryptedText);
    let remainingAttempts = 3;
    const attempt = () => {
      const result = pending.copy();
      if (!result && --remainingAttempts) {
        setTimeout(attempt);
      } else {
        pending.destroy();
      }
    };
    attempt();
  }
}
