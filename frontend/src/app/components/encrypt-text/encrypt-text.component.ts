import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Clipboard } from '@angular/cdk/clipboard';


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
  constructor(private clipboard: Clipboard) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.notesToEncrypt.valid) {
      const text = this.notesToEncrypt.get('text')?.value
      const passphrase = this.notesToEncrypt.get('passphrase')?.value
      console.log(text)
      console.log(passphrase)
      this.encryptedText = text + "hey"
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
