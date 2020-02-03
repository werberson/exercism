function isStopCodon(codon) {
  return codon === 'UAA' || codon === 'UAG' || codon === 'UGA';
}

const protein = (codon) => {
  switch (codon) {
    case 'AUG':
      return 'Methionine';
    case 'UUU':
    case 'UUC':
      return 'Phenylalanine';
    case 'UUA':
    case 'UUG':
      return 'Leucine';
    case 'UCU':
    case 'UCC':
    case 'UCA':
    case 'UCG':
      return 'Serine';
    case 'UAU':
    case 'UAC':
      return 'Tyrosine';
    case 'UGU':
    case 'UGC':
      return 'Cysteine';
    case 'UGG':
      return 'Tryptophan';
    default:
      throw new Error('Invalid codon');
  }
};

export const translate = (rna = '') => {
  const codonSize = 3;
  if (rna.length % codonSize !== 0) {
        throw new Error('Invalid codon');
    }
    const proteins = [];
    for (let i = 0; i < rna.length; i = i + codonSize) {
        const codon = rna.substr(i, codonSize);
        if (isStopCodon(codon)) {
          break
        }
        proteins.push(protein(codon));
    }
    return proteins;
};