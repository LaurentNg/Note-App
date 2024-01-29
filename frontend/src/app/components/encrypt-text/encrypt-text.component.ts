import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Clipboard } from '@angular/cdk/clipboard';
import { EncryptorHttpService } from 'src/app/services/http/encryptor-http.service';
import { Encryptor } from 'src/app/models/interfaces/encryptor';


@Component({
  selector: 'app-encrypt-text',
  templateUrl: './encrypt-text.component.html',
  styleUrl: './encrypt-text.component.scss'
})
export class EncryptTextComponent implements OnInit {
  notesToEncrypt = new FormGroup({
    text: new FormControl<string>('', [Validators.required]),
    passphrase: new FormControl<string>('', [Validators.required])
  });
  encryptedText = ""
  constructor(private clipboard: Clipboard, private encryptor: EncryptorHttpService) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.notesToEncrypt.valid) {
      const text = this.notesToEncrypt.get('text')?.value
      const passphrase = this.notesToEncrypt.get('passphrase')?.value   
      const encryptorObj = new Encryptor(text as string, passphrase as string)
      this.encryptor.encrypt(encryptorObj).subscribe((res) => {
        if ('text' in res) this.encryptedText = res.text
      })
    }
    else console.log('not valid')
  }

  copyEncryptedText() {
    const pending = this.clipboard.beginCopy(this.encryptedText);
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
